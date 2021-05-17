import { matchPath } from 'react-router';

/**
 * Authentication status constants.
 */
export enum AuthStatus {
  UNKNOWN,
  SUCCESS,
  FAILURE,
}

/**
 * ...
 */
export function AccessibleRouteChangeHandler() {
  return window.setTimeout(() => {
    const mainContainer = document.getElementById('primary-app-container');
    if (mainContainer) {
      mainContainer.focus();
    }
  }, 50);
}

/**
 * Returns date time in readable format
 * @param timestamp Current date time value.
 * @returns Formatted date time.
 */
export function GetReadableTimestamp(timestamp: Date): string {
  let readable = ""
  readable += GetReadableDate(timestamp) + " "
  readable += GetReadableTime(timestamp)
  return readable
}

/**
 * Returns date in readable format
 * @param timestamp Current date time value.
 * @returns Formatted date.
 */
export function GetReadableDate(timestamp: Date): string {
  let readable = ""
  readable += timestamp.getDate() + ". "
  readable += timestamp.getMonth() + 1 + ". "
  readable += timestamp.getFullYear()
  return readable
}

/**
 * Returns time in readable format
 * @param timestamp Current date time value.
 * @returns Formatted time.
 */
export function GetReadableTime(timestamp: Date): string {
  let readable = ""
  readable += ("0" + timestamp.getHours()).slice(-2) + ":"
  readable += ("0" + timestamp.getMinutes()).slice(-2) + ":"
  readable += ("0" + timestamp.getSeconds()).slice(-2)
  return readable
}

/**
 * UserStatusBoolToStringConvert converts bool on input to readable string. It
 * determines if account is active or inactive.
 * 
 * @param enabled Account status
 */
export function UserStatusBoolToStringConvert(enabled: boolean): string {
  if (enabled === true) {
    return "Aktivní";
  }
  return "Deaktivován";
}

/**
 * GetIDByPath searching for URL parameter ID. It requires pathname (string
 * which we are looking for) and pathPattern which is whole URL. The function
 * returns string as parameter from URL.
 * 
 * @param pathname Matching string
 * @param pathPattern Full path
 */
export function GetIDByPath(pathname, pathPattern): string {
  const match = matchPath(pathname, {
    path: pathPattern,
    exact: true,
    strict: false
  })

  return match.params.id
}
