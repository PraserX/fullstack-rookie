import { ErrorForbidden, ErrorNotFound, IResponse } from './ApiHandler'

/**
 * Error handling function. 
 * @param fatal Fatal error indication. If true then error page will show up.
 * @param response Response with error inside. If status code is 500, then
 * error page will show up. Otherwise error message will be placed in console.
 * @param callback Redirect handler function callback.
 */
export const HandleError = (fatal: boolean, response: IResponse | null, callback: (path: string) => void): void => {
  if (fatal || (response === null)) {
    callback("/error/connection-lost")
  } else if (response.error === ErrorNotFound) {
    console.error("Error: " + response.error)
    callback("/error/not-found")
  } else if (response.error === ErrorForbidden) {
    console.error("Error: " + response.error)
  } else {
    console.error("Error: " + response.error)
    callback("/error")
  }
}