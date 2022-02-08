import { client } from './client'

/**
 * LoginUserParam type.
 */
export type LoginUserParam = {
  name: string
  password: string
}

/**
 * LoginUserRes type.
 */
export type LoginUserRes = {
  jwt: string
}

/**
 * request user login.
 *
 * @param param user login data.
 */
export const loginUser = async (
  param: LoginUserParam
): Promise<LoginUserRes> => {
  const res = await client.post<LoginUserRes>('/login', param)

  return res.data
}
