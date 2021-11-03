import Axios from "axios";
import React from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Nav from "./components/Nav";
import Legal from "./components/Legal";

import ProtectedRoute from "./components/ProtectedRoute";
import Applications from "./pages/Applications";
import Browse from "./pages/Browse";
import LandingPage from "./pages/LandingPage";
import OAuthCallback from "./src/OAuthCallback";
import LoginPage from "./pages/LoginPage";

Axios.defaults.baseURL = process.env.BASE_URL;
Axios.defaults.withCredentials = true;

const App: React.FC = () => {
  return (
<<<<<<< Updated upstream
    <>
      <Nav />
      <div style={{minHeight: '100%'}}>
        <Router>
          <Switch>
            {/* Oauth callback */}
            <Route path="/oauth/callback/google" component={OAuthCallback} />
            {/* Protected routes */}
            <ProtectedRoute>
              {/* Landing page */}
              <Route exact path="/" component={LandingPage} />
              {/* Browse page */}
              <Route exact path="/browse" component={Browse} />
              {/* Applications page */}
              <Route exact path="/applications" component={Applications} />
            </ProtectedRoute>
          </Switch>
        </Router>
      </div>
      <Legal />
    </>
=======
    <Router>
      <Switch>
        {/* Oauth callback */}
        <Route path="/oauth/callback/google" component={OAuthCallback} />
        <Route exact path="/login" component={LoginPage}/>
        {/* Protected routes */}
        <ProtectedRoute>
          {/* Landing page */}
          <Route exact path="/" component={LandingPage} />
          {/* Browse page */}
          <Route exact path="/browse" component={Browse} />
          {/* Applications page */}
          <Route exact path="/applications" component={Applications} />
        </ProtectedRoute>
      </Switch>
    </Router>
>>>>>>> Stashed changes
  );
};

export default App;
