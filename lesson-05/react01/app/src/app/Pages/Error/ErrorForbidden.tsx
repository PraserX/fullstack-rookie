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

const ErrorForbidden: React.FunctionComponent = () => (
  <DefaultAppLayout user={undefined}>
    <PageSection>
      <EmptyState>
        <EmptyStateIcon icon={ExclamationTriangleIcon} />
        <Title headingLevel="h4" size="lg">
          Přístup odepřen
        </Title>
        <EmptyStateBody>
          Přístup na tuto stránku je Vám odepřen. Pokud se domníváte, že jedná o chybu, kontaktuje administrátora.
        </EmptyStateBody>
        <br /><br />
        <NavLink to="/" className="pf-c-button">Zpět na úvodní stránku</NavLink>
      </EmptyState>
    </PageSection>
  </DefaultAppLayout>
)

export { ErrorForbidden };
