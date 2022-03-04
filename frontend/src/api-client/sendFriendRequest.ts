import { client } from './client'

/**
 * SendFriendRequestParams type.
 */
export type SendFriendRequestParams = {
  myname: string
  othername: string
}

/**
 * send friend request.
 *
 * @param myname my name.
 * @param othername other name.
 */
export const sendFriendRequest = async (params: SendFriendRequestParams) => {
  const { data } = await client.post('/api/friend/request', params)

  console.log(data)
}
