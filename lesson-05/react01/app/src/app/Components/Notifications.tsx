import * as React from 'react';
import {
  Alert,
  AlertGroup,
  AlertActionCloseButton
} from '@patternfly/react-core';

export interface INotificationContent {
  variant: "default" | "success" | "warning" | "danger" | "info" | undefined;
  title: string;
}

export interface FCNotificationsProps {
  children?: React.ReactNode;
  notifications: Array<INotificationContent>;
  setNotificationsCallback: any;
}

export const FCNotifications: React.FunctionComponent<FCNotificationsProps> = (
  { children = null, notifications = Array(), setNotificationsCallback, ...props }: FCNotificationsProps) => {

  return (
    <AlertGroup isToast>
      {notifications.map(({ variant, title }, idx) => (
        <Alert
          isLiveRegion
          variant={variant}
          title={title}
          timeout={5000}
          actionClose={<AlertActionCloseButton onClose={() => setNotificationsCallback(notifications.splice(1))} />} />

      ))}
    </AlertGroup>
  );
}