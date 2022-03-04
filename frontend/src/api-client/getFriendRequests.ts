import { client } from './client'

/**
 * get friend requests
 */
export const getFriendRequests = async () => {
  const { data } = await client.get('/api/getfriendrequest')
  console.log(data)
}
