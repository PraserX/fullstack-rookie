import * as React from 'react';
import { NavLink } from 'react-router-dom';
import {
  EmptyState,
  EmptyStateBody,
  EmptyStateIcon,
  Title,
  PageSection
} from '@patternfly/react-core';
import ExclamationCircleIcon from '@patternfly/react-icons/dist/js/icons/exclamation-circle-icon';

export const ErrorInternal: React.FunctionComponent = () => (
  <PageSection>
    <EmptyState>
      <EmptyStateIcon icon={ExclamationCircleIcon} />
      <Title headingLevel="h4" size="lg">
        Něco se pokazilo
        </Title>
      <EmptyStateBody>
        Na severu došlo k chybě a nebylo možné operaci dokončit. Zkuste to prosím později. Pokud chyba přetrvává, kontaktujte prosím administrátora.
      </EmptyStateBody>
      <br /><br />
      <NavLink to="/" className="pf-c-button">Zpět na úvodní stránku</NavLink>
    </EmptyState>
  </PageSection>
)
