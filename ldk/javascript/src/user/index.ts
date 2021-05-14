import {promisify} from "../promisify";

export interface User {
    requestJWT(): Promise<string>;
}

export function requestJWT(): Promise<string> {
    return promisify(oliveHelps.user.requestJWT)
}
