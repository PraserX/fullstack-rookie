import * as React from 'react';
import { NavLink, Link } from 'react-router-dom';
import {
  Nav,
  Banner,
  NavList,
  NavItem,
  Page,
  PageHeader,
  SkipToContent,
  Dropdown,
  DropdownGroup,
  DropdownToggle,
  DropdownItem,
  DropdownSeparator,
  PageHeaderTools,
  PageHeaderToolsGroup,
  PageHeaderToolsGroupProps,
  PageHeaderToolsItem,
} from '@patternfly/react-core';
import { routes } from '@app/routes';

interface IAppLayout {
  children: React.ReactNode;
}

const DefaultAppLayout: React.FunctionComponent<IAppLayout> = ({ children }) => {
  const logoProps = {
    href: '/',
    target: '_blank'
  };

  const defaultVisibility: PageHeaderToolsGroupProps = {
    children: {},
    visibility: {
      default: "hidden",
      sm: "hidden",
      md: "hidden",
      lg: "visible",
      xl: "visible",
      '2xl': "visible",
    }
  }

  const [isDropdownOpen, setIsDropdownOpen] = React.useState(false);

  const onDropdownSelect = () => {
    setIsDropdownOpen(false);
  }

  const onDropdownToggle = () => {
    if (isDropdownOpen) {
      setIsDropdownOpen(false);
    } else {
      setIsDropdownOpen(true);
    }
  }

  const headerTools = (
    <PageHeaderTools>
      <PageHeaderToolsGroup>
        <PageHeaderToolsItem visibility={defaultVisibility.visibility}>

        </PageHeaderToolsItem>
      </PageHeaderToolsGroup>
    </PageHeaderTools>
  )

  const Navigation = (
    <Nav id="nav-primary-simple" theme="dark" variant="horizontal">
      <NavList id="nav-list-simple">
        {routes.map((route, idx) => route.label && route.navigation === true && (
          <NavItem key={`${route.label}-${idx}`} id={`${route.label}-${idx}`}>
            <NavLink exact to={route.path} activeClassName="pf-m-current">{route.label}</NavLink>
          </NavItem>
        ))}
      </NavList>
    </Nav>
  );

  const Header = (
    <PageHeader
      logo="My Comments"
      logoProps={logoProps}
      headerTools={headerTools}
      topNav={Navigation}
    />
  )

  const PageSkipToContent = (
    <SkipToContent href="#primary-app-container">
      Continue
    </SkipToContent>
  );

  return (
    <React.Fragment>
      <Page
        mainContainerId="primary-app-container"
        header={Header}
        skipToContent={PageSkipToContent}>
        {children}
      </Page>
    </React.Fragment>
  );
}

export { DefaultAppLayout };
