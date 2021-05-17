import { ICommentResponse } from '@app/Types/Responses/Comment'
import { IUserResponse } from '@app/Types/Responses/User'

export const ErrorForbidden = "forbidden"
export const ErrorNotFound = "not found"
export const ErrorUnwantedRedirect = "unwanted dangerous redirect"
export const ErrorNetwork = "general unknown network error"
export const ErrorUnexpected = "unexpected error"

export interface IResponse {
  // Response state (true == success, false == failure).
  ok: boolean;
  // Raw JSON response promise.
  raw?: Promise<any>;
  // Parsed JSON data.
  data?: any;
  // Error message.
  error?: string;
  // Raw original response (returned alongside error message).
  response?: Response;
}

const GetFetchParams = (requestMethod: string, csrfToken: string, requestBody: string): RequestInit => {
  return {
    method: requestMethod, // *GET, POST, PUT, DELETE, etc.
    mode: 'same-origin', // no-cors, *cors, same-origin
    cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
    credentials: 'same-origin', // include, *same-origin, omit
    headers: {
      'Content-Type': 'application/json',
      'X-CSRF-Token': csrfToken,
    },
    redirect: 'follow', // manual, *follow, error
    referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
    body: requestBody // body data type must match "Content-Type" header
  }
}

const GetPromise = (response: Response): Promise<IResponse> => {
  return new Promise<IResponse>(function (resolve, reject) {
    if (response.ok) {
      resolve({ ok: true, raw: response.json() } as IResponse)
    } else if (response.status === 403) {
      reject({ ok: false, error: ErrorForbidden, response: response } as IResponse)
    } else if (response.status === 404) {
      reject({ ok: false, error: ErrorNotFound, response: response } as IResponse)
    } else if (response.status >= 500) {
      reject({ ok: false, error: ErrorUnexpected, response: response } as IResponse)
    } else if (response.type === "opaqueredirect") {
      reject({ ok: false, error: ErrorUnwantedRedirect, response: response } as IResponse)
    } else {
      reject({ ok: false, error: ErrorNetwork, response: response } as IResponse)
    }
  })
}

/**
 * Generic API handler. This method is quite specific. It tries to fetch data
 * three times due to authorization proxy, which could be between user and
 * server. Due to authorization proxy, some request may end up unsuccessful
 * because of access token expiration. This is why we have to try to get
 * response multiple times.
 * 
 * @param url Request URI promise.
 * @param method Request method.
 * @param request Request CSRF token.
 * @param requestBody Request body.
 */
export const APIRequest = async (url: string, method: string, csrfToken: string = "", requestBody: string = ""): Promise<IResponse> => {
  let response: Response = {} as Response

  let fetchParams = method == 'GET'
    ? { method: 'GET', redirect: 'follow' } as RequestInit
    : GetFetchParams(method, csrfToken, requestBody)

  for (let i = 0; i < 3; i++) {
    response = await fetch(url, fetchParams)
      .catch(error => new Response(JSON.stringify({ ok: false, statusText: "unexpected fatal error: " + error })))

    if (!response.redirected && ((response.status < 300) || (response.status > 499))) {
      break
    }
  }

  return GetPromise(response)
}

export const APIGetRequest = async (url: string): Promise<IResponse> => {
  return APIRequest(url, 'GET')
}

export const APIPostRequest = async (url: string, csrfToken: string, requestBody: string): Promise<IResponse> => {
  return APIRequest(url, 'POST', csrfToken, requestBody)
}

export const APIPutRequest = async (url: string, csrfToken: string, requestBody: string): Promise<IResponse> => {
  return APIRequest(url, 'PUT', csrfToken, requestBody)
}

export const APIDeleteRequest = async (url: string, csrfToken: string): Promise<IResponse> => {
  return APIRequest(url, 'DELETE', csrfToken, "")
}

export const APIGetComments = async (setterCallback: any, comments: any): Promise<IResponse> => {
  return new Promise<IResponse>(function (resolve, reject) {
    APIGetRequest("/api/v1/comments")
      .then(response => {
        if (response.ok) {
          return response.raw
        }
        throw { ok: false, error: ErrorUnexpected, response: response } as IResponse
      })
      .then(data => {
        if (data !== null && data !== undefined) {
          if ((data.comments === null) || (data.comments === undefined)) {
            data.comments = Array<ICommentResponse>()
          }
          if (JSON.stringify(data.comments) !== JSON.stringify(comments)) {
            setterCallback(data.comments.map(item => item as ICommentResponse))
          }
          resolve({ ok: true, data: data } as IResponse)
        }
        resolve({ ok: false } as IResponse)
      })
      .catch(error => reject(error as IResponse))
  })
}

/**
 * Method provides acquisition of REST API newsletters data. It returns list of newsletters. 
 * 
 * @param setterCallback Standard React setter.
 * @param newsletters React object.
 */
export const APIGetCSRF = async (url: string, setterCSRFCallback: (any | null)): Promise<IResponse> => {
  return new Promise<IResponse>(function (resolve, reject) {
    APIGetRequest(url)
      .then(response => {
        if (response.ok) {
          return response.raw
        }
        throw { ok: false, error: ErrorUnexpected, response: response } as IResponse
      })
      .then(data => {
        if (data !== null && data !== undefined) {
          if (setterCSRFCallback !== null) {
            setterCSRFCallback(data.csrf)
          }
          resolve({ ok: true, data: data } as IResponse)
        }
        resolve({ ok: false } as IResponse)
      })
      .catch(error => reject(error as IResponse))
  })
}