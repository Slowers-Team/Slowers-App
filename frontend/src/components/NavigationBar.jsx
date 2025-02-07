import "../App.css";
import { useTranslation } from "react-i18next";
import { Navbar, Nav, NavDropdown, Button, Offcanvas } from "react-bootstrap";
import LangSelect from "./LangSelect";
import { Link, useLoaderData, Outlet, useFetcher } from "react-router-dom";
import { useState } from "react";
import { Authenticator } from "../Authenticator";

export const NavigationBar = () => {
  const { t, i18n } = useTranslation();
  const [showOffCanvas, setShowOffCanvas] = useState(false);
  const { isLoggedIn, username } = useLoaderData();
  const fetcher = useFetcher();

  const handleClose = () => setShowOffCanvas(false);
  const handleShow = () => setShowOffCanvas(!showOffCanvas);

  const handleLogout = () =>
    fetcher.submit({}, { action: "/logout", method: "post" });

  return (
    <>
      <Navbar expand="sm" className="nav-bar">
        <Button id="offcanvasButton" variant="light" className="menu-button mx-2" onClick={handleShow}>
          <span className="navbar-toggler-icon"></span>
        </Button>
        <Navbar.Brand as={Link} to="/">
          <h1>Slowers</h1>
        </Navbar.Brand>
        <Nav className="ms-auto mx-2">
          {isLoggedIn && (
            <NavDropdown title={username} id="collasible-nav-dropdown" align="end">
              <Nav.Link className="text-secondary" as={Link} to="/user">
                <i className="bi bi-person-circle"> </i>
                {t("menu.profile")}
              </Nav.Link>
              <Nav.Link
                className="text-secondary"
                as={Link}
                onClick={() => {
                  handleLogout();
                  handleClose();
                }}
              >
                <i className="bi bi-door-open-fill"> </i>
                {t("menu.logout")}
              </Nav.Link>
            </NavDropdown>
          )}
          <NavDropdown title={<i className="bi bi-globe-americas"></i>} id="languageButton" align="end" className="lang-menu">
            <LangSelect />
          </NavDropdown>
        </Nav>
      </Navbar>

      <Offcanvas
        show={showOffCanvas}
        onHide={handleClose}
        className="offcanvas-thin"
      >
        <Offcanvas.Header closeButton>
          <Offcanvas.Title>
            <h2>Slowers</h2>
          </Offcanvas.Title>
        </Offcanvas.Header>
        <Offcanvas.Body>
          <Nav className="flex-column pe-3">
            <Nav.Link
              className="text-secondary"
              as={Link}
              to="/"
              onClick={handleClose}
            >
              <i className="bi bi-house"> </i>
              {t("menu.home")}
            </Nav.Link>
            {!isLoggedIn && (
              <Nav.Link
                className="text-secondary"
                as={Link}
                to="/login"
                onClick={handleClose}
              >
                <i className="bi bi-box-arrow-in-right"> </i>
                {t("menu.login")}
              </Nav.Link>
            )}
            {!isLoggedIn && (
              <Nav.Link
                className="text-secondary"
                as={Link}
                to="/register"
                onClick={handleClose}
              >
                <i className="bi bi-person-add"> </i>
                {t("menu.register")}
              </Nav.Link>
            )}
            {isLoggedIn && (
              <Nav.Link
                className="text-secondary"
                as={Link}
                to="/retailer"
                onClick={handleClose}
              >
                <i className="bi bi-shop-window"> </i>
                {t("menu.retailer")}
              </Nav.Link>
            )}
            {isLoggedIn && (
              <Nav.Link
                className="text-secondary"
                as={Link}
                to="/grower"
                onClick={handleClose}
              >
                <i className="bi bi-flower1"> </i>
                {t("menu.grower")}
              </Nav.Link>
            )}
            {isLoggedIn && (
              <Nav.Link
                className="text-secondary"
                as={Link}
                to="/business_owner"
                onClick={handleClose}
              >
                <i className="bi bi-flower1"> </i>
                {t("menu.businessowner")}
              </Nav.Link>
            )}
            <Nav.Link
              className="text-secondary"
              as={Link}
              to="/terms"
              onClick={handleClose}
            >
              <i className="bi bi-file-earmark-text"> </i>
              {t("menu.terms")}
            </Nav.Link>
          </Nav>
        </Offcanvas.Body>
      </Offcanvas>
      <Outlet />
    </>
  );
};

export default NavigationBar;
