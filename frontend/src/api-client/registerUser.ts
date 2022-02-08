import { client } from './client'

/**
 * RegisterUserParam type.
 */
export type RegisterUserParam = {
  name: string
  email: string
  password: string
}

/**
 * RegisterUserRes type.
 */
export type RegisterUserRes = {
  jwt: string
}

/**
 * request user registration.
 *
 * @param param user registeration data.
 */
export const registerUser = async (
  param: RegisterUserParam
): Promise<RegisterUserRes> => {
  const res = await client.post<RegisterUserRes>('/register', {
    ...param,
  })

  return res.data
}
