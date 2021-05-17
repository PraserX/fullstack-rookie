import * as React from 'react';

import { Route, RouteComponentProps, Switch } from 'react-router-dom';
import { AccessibleRouteChangeHandler } from '@app/Utils/Helpers';
import { UseDocumentTitle } from '@app/Utils/UseDocumentTitle';
import { LastLocationProvider, useLastLocation } from 'react-router-last-location';

import { DefaultAppLayout } from '@app/Layouts/AppLayout';
import { AuthStatus } from '@app/Utils/Helpers';

import { Dashboard } from '@app/Pages/Dashboard/Dashboard';
import { ErrorInternal } from '@app/Pages/Error/ErrorInternal';
import { ErrorConnection } from '@app/Pages/Error/ErrorConnection';
import { ErrorNotFound } from '@app/Pages/Error/ErrorNotFound';
import { IUserResponse } from './Types/Responses/User';
import { ErrorForbidden } from './Pages/Error/ErrorForbidden';

let routeFocusTimer: number;

export interface IAppRoute {
  label?: string;
  /* eslint-disable @typescript-eslint/no-explicit-any */
  component: React.ComponentType<RouteComponentProps<any>> | React.ComponentType<any>;
  /* eslint-enable @typescript-eslint/no-explicit-any */
  exact?: boolean;
  path: string;
  title: string;
  navigation: boolean;
  isAsync?: boolean;
  authenticated?: AuthStatus | undefined;
  currentUser?: IUserResponse;
  version?: string;
}

const routes: IAppRoute[] = [
  {
    component: Dashboard,
    exact: true,
    label: 'Přehled',
    path: '/',
    navigation: true,
    title: 'MCA | Comments',
  },
  {
    component: ErrorInternal,
    exact: true,
    isAsync: true,
    label: 'System error',
    path: '/error',
    navigation: false,
    title: 'MCA | System error',
  },
  {
    component: ErrorForbidden,
    exact: true,
    isAsync: true,
    label: 'Access denied',
    path: '/error/forbidden',
    navigation: false,
    title: 'MCA | Access denied',
  },
  {
    component: ErrorNotFound,
    exact: true,
    isAsync: true,
    label: 'Page not found',
    path: '/error/not-found',
    navigation: false,
    title: 'MCA | Page not found',
  },
  {
    component: ErrorConnection,
    exact: true,
    isAsync: true,
    label: 'Connection lost',
    path: '/error/connection-lost',
    navigation: false,
    title: 'MCA | Connection lost',
  },
];

const useA11yRouteChange = (isAsync: boolean) => {
  const lastNavigation = useLastLocation();
  React.useEffect(() => {
    if (!isAsync && lastNavigation !== null) {
      routeFocusTimer = AccessibleRouteChangeHandler();
    }
    return () => {
      window.clearTimeout(routeFocusTimer);
    };
  }, [isAsync, lastNavigation]);
};

const RouteWithTitleUpdates = ({ component: Component, isAsync = false, title, authenticated = undefined, currentUser = undefined, version = undefined, ...rest }: IAppRoute) => {
  useA11yRouteChange(isAsync);
  UseDocumentTitle(title);

  function routeWithTitle(routeProps: RouteComponentProps) {
    return <DefaultAppLayout><Component {...rest} {...routeProps} /></DefaultAppLayout>;
  }

  return <Route render={routeWithTitle} />;
};

const PageNotFound = ({ title }: { title: string }) => {
  UseDocumentTitle(title);
  return <Route component={ErrorNotFound} />;
};

const AppRoutes = (): React.ReactElement => {
  return (
    <LastLocationProvider>
      <Switch>
        {routes.map(({ path, exact, component, title, isAsync, navigation }, idx) => (
          <RouteWithTitleUpdates
            path={path}
            exact={exact}
            component={component}
            key={idx}
            title={title}
            navigation={navigation}
            isAsync={isAsync}
          />
        ))}
        <PageNotFound title="404 Page Not Found!" />
      </Switch>
    </LastLocationProvider >
  )
};

export { AppRoutes, routes };
