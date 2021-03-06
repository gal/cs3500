import React from 'react';
import {Route, RouteProps} from 'react-router-dom';

import {User} from '../api/types';
import EditProfileForm from './EditProfileForm';
import {LoginModal} from './LoginModal';

type ProtectedRouteProps = RouteProps & {
  user?: User;
  children: React.ReactNode;
};

const ProtectedRoute = ({user, children, ...rest}: ProtectedRouteProps) => {
  console.log(user);
  return (
    <Route
      {...rest}
      render={({location}) =>
        user ? (
          user.username && user.tags ? (
            children
          ) : (
            <EditProfileForm type="modal" />
          )
        ) : (
          <LoginModal />
        )
      }
    />
  );
};

export default ProtectedRoute;
