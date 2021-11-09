import Axios from 'axios';
import ProjectsPage from './pages/ProjectsPage';
import React from 'react';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';

import Legal from './components/Legal';
import NavBar from './components/Nav';
import ProtectedRoute from './components/ProtectedRoute';
import Applications from './pages/Applications';
import Browse from './pages/Browse';
import LandingPage from './pages/LandingPage';
import OAuthCallback from './src/OAuthCallback';

Axios.defaults.baseURL = process.env.API_URL;
Axios.defaults.withCredentials = true;

const App: React.FC = () => {
  return (
    <>
      <Router>
        <NavBar />
        <Switch>
          {/* Oauth callback */}
          <Route path="/oauth/callback/google" component={OAuthCallback} />
          {/* Protected routes */}
          <ProtectedRoute>
            {/* Landing page */}
            <Route exact path="/" component={LandingPage} />
            {/* Browse page */}
            <Route exact path="/browse" component={Browse} />
            {/* Projects page */}
            <Route exact path="/projects" component={ProjectsPage} />
            {/* Applications page */}
            <Route exact path="/applications" component={Applications} />
          </ProtectedRoute>
        </Switch>
        <Legal />
      </Router>
    </>
  );
};

export default App;
