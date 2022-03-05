import { Diary } from '@/types/Diary'
import { client } from './client'

/**
 * diary delete.
 *
 * @param params params.
 */
export const diaryDelete = async (params: Diary): Promise<Diary> => {
  const { data } = await client.post<Diary>('/api/diary/delete', params)

  return data
}
