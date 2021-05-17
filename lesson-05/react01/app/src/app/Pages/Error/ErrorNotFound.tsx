import * as React from 'react';
import { NavLink } from 'react-router-dom';
import {
  EmptyState,
  EmptyStateBody,
  EmptyStateIcon,
  Title,
  PageSection
} from '@patternfly/react-core';
import ExclamationTriangleIcon from '@patternfly/react-icons/dist/js/icons/exclamation-triangle-icon';
import { DefaultAppLayout } from '@app/Layouts/AppLayout';

const ErrorNotFound: React.FunctionComponent = () => (
  <PageSection>
    <EmptyState>
      <EmptyStateIcon icon={ExclamationTriangleIcon} />
      <Title headingLevel="h4" size="lg">
        Stránka nenalezena
        </Title>
      <EmptyStateBody>
        Vámi hledaná stránka nebyla nalezena. Jedná se o tzv. chybový stav 404, což znamená, že přistupujete na stránku, která neexistuje.
        </EmptyStateBody>
      <br /><br />
      <NavLink to="/" className="pf-c-button">Zpět na úvodní stránku</NavLink>
    </EmptyState>
  </PageSection>
)

export { ErrorNotFound };