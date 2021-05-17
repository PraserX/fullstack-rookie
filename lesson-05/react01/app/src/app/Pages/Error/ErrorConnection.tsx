import * as React from 'react';
import { NavLink } from 'react-router-dom';
import {
  EmptyState,
  EmptyStateBody,
  EmptyStateIcon,
  Title,
  PageSection
} from '@patternfly/react-core';
import DisconnectedIcon from '@patternfly/react-icons/dist/js/icons/disconnected-icon';

export const ErrorConnection: React.FunctionComponent = () => (
  <PageSection>
    <EmptyState>
      <EmptyStateIcon icon={DisconnectedIcon} />
      <Title headingLevel="h4" size="lg">
        Spojení se servem je ztraceno
        </Title>
      <EmptyStateBody>
        Momentálně se není možné spojit se serverem. Zkuste to prosím později. Pokud chyba přetrvává, kontaktujte prosím administrátora.
      </EmptyStateBody>
    </EmptyState>
  </PageSection>
)
