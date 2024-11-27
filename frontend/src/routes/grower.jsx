import GrowerSitesPage from "../pages/GrowerSitesPage"
import GrowerHomePage from "../pages/GrowerHomePage" 
import GrowerFlowerPage from "../pages/GrowerFlowerPage" 
import GrowerImagesPage from "../pages/GrowerImagesPage" 
import { rootFlowerLoader, rootSiteLoader, siteFlowerLoader, siteLoader } from "../loaders/grower"

export default growerRoutes = { 
  path: "grower",
  element: <GrowerLayout />,
  async action({ request}) {
    const formData = await request.formData()
    return redirect(formData.get("redirect")) // redirect user after site deletion
  },
  children: [
    { index: true,
      element: <GrowerHomePage />,
      loader() {return false}
    },
    { path: "flowers", 
      element: <GrowerFlowerPage />,
      loader: rootFlowerLoader
    },
    { path: "sites",
      element: <GrowerSitesPage />,
      loader: rootSiteLoader,
      action() {return null} // we only want to reload data
    }, 
    { path: ":siteId",
      id: "site",
      loader: siteLoader, 
      children: [
        { index: true,
          element: <GrowerHomePage />
        },
        { path: "flowers",
          element: <GrowerFlowerPage />,
          loader: siteFlowerLoader
        },
        { path: "sites",
          element: <GrowerSitesPage />
        },
        { path: "images",
          element: <GrowerImagesPage />
        }
      ] } 
  ] 
}
