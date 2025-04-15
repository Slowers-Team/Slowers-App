import {
  createBrowserRouter,
  RouterProvider,
  redirect,
} from "react-router-dom";
import { useEffect } from "react";
import { useTranslation } from "react-i18next";
import NavigationBar from "./components/NavigationBar";

import { Authenticator } from "./Authenticator";
import 'bootstrap-icons/font/bootstrap-icons.css';

import HomeLayout from "./layouts/HomeLayout";
import GrowerLayout from "./layouts/GrowerLayout";
import MarketplaceLayout from "./layouts/MarketplaceLayout";
import BusinessLayout from "./layouts/BusinessLayout";
import UserPageLayout from "./layouts/UserPageLayout";

import HomePage from "./pages/HomePage";
import RegisterPage from "./pages/RegisterPage";
import TermsPage from "./pages/TermsPage";
import LogInPage from "./pages/LogInPage";
import UserPage from "./pages/UserPage";
import RetailerHomePage from "./pages/RetailerHomePage";
import RetailerFlowerPage from "./pages/RetailerFlowerPage";
import MarketplaceHomePage from "./pages/marketplaceHomePage";
import GrowerHomePage from "./pages/GrowerHomePage";
import GrowerFlowerPage from "./pages/GrowerFlowerPage";
import GrowerSitesPage from "./pages/GrowerSitesPage";
import GrowerImagesPage from "./pages/GrowerImagesPage";
import BusinessPage from "./pages/BusinessPage";
import BusinessEmployeesPage from "./pages/BusinessEmployeesPage";


const Root = () => {
  const { t, i18n } = useTranslation();

  useEffect(() => {
    setLanguage();
  }, []);

  const setLanguage = () => {
    const langCookie = document.cookie
      .split("; ")
      .find((row) => row.startsWith("lang="));
    const language = langCookie ? langCookie.split("=")[1] : "en";
    i18n.changeLanguage(language);
  };

  return (
    <div>
      <NavigationBar />
    </div>
  );
};

// Redirect user to a default role, if they are logged in
const roleLoader = () => {
  if (Authenticator.isLoggedIn) {
    if (Authenticator.role === "grower") {
      return redirect("/grower");
    } else {
      return redirect("/retailer");
    }
  }
  return null;
};

// Redirect user to login-screen, if they are not logged in
function protectedLoader() {
  if (!Authenticator.isLoggedIn) {
    return redirect("/login");
  }
  return null;
}

// Redirect user to home page if they are not allowed to access the page
function authorizeAccess() {
  if (!Authenticator.isLoggedIn) {
    return redirect("/login");
  }
  const path = window.location.pathname

  if (path.startsWith("/grower") && !( Authenticator.businessType === "grower" && (Authenticator.designation === "owner" || Authenticator.designation === "employee"))) {
    return redirect("/home")
  }
  if (path.startsWith("/business/employees") && !( Authenticator.designation === "owner" || Authenticator.designation === "employee" )) {
    return redirect("/home")
  }
  if (path.startsWith("/business/retailer") && !( Authenticator.businessType === "retailer" && (Authenticator.designation === "owner" || Authenticator.designation === "employee"))) {
    return redirect("/home")
  }
  return null;
}

// Redirect user to home page if logged in, else to login
const rootRedirect = () => {
  if (Authenticator.isLoggedIn) {
    return redirect("/home")
  } else {
    return redirect("/login");
  }
};

const rootLoader = () => {
  Authenticator.refresh(); // try to fetch login info from local storage
  return {
    role: Authenticator.role,
    isLoggedIn: Authenticator.isLoggedIn,
    username: Authenticator.username,
  };
};

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    id: "root",
    loader: rootLoader,
    children: [
      {
        index: true,
        loader: rootRedirect, // rootLoader always redirects to another place
      },
      {
        path: "login",
        loader: roleLoader,
        element: <LogInPage />,
        action() {
          return redirect("/");
        }, // POST /login -> redirect to homepage
      },
      {
        path: "register",
        loader: roleLoader,
        element: <RegisterPage />,
      },
      { path: "terms", element: <TermsPage /> },
      {
        path: "logout",
        action() {
          return Authenticator.logout();
        }, // POST /logout -> Authenticator.logout()
      },
      {
        path: "*",
        loader: protectedLoader,
        children: [
          {
            path: "home",
            element: <HomeLayout />,
            children: [
              { index: true, element: <HomePage />}
            ],
          },
          {
            path: "grower",
            loader: authorizeAccess,
            element: <GrowerLayout />,
            children: [
              { index: true, element: <GrowerHomePage /> },
              { path: "flowers", element: <GrowerFlowerPage /> },
              { path: "sites", element: <GrowerSitesPage /> },
              {
                path: ":siteId",
                children: [
                  { index: true, element: <GrowerHomePage /> },
                  { path: "flowers", element: <GrowerFlowerPage /> },
                  { path: "sites", element: <GrowerSitesPage /> },
                  { path: "images", element: <GrowerImagesPage /> },
                ],
              },
            ],
          },
          {
            path: "marketplace",
            loader: authorizeAccess,
            element: <MarketplaceLayout />,
            children: [
              { index: true, element: <MarketplaceHomePage /> },
              { path: "flowers", element: <RetailerFlowerPage /> },
            ],
          },
          {
            path: "business",
            loader: authorizeAccess,
            element: <BusinessLayout />,
            children: [
              { index: true, element: <BusinessPage /> },
              { path: "employees", element: <BusinessEmployeesPage /> },
              { path: "retailer", element: <RetailerHomePage /> },
            ]
          },
          { 
            path: "user",
            loader: authorizeAccess,
            element: <UserPageLayout />,
            children: [
              { index: true, element: <UserPage /> }
            ]
          },
          {
            path: "*",
            loader() {
              return redirect("/");
            },
          }, // redirect undefined paths to home
        ],
      },
    ],
  },
]);

export default function App() {
  return <RouterProvider router={router} />;
}
