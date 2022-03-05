import { Diary } from '@/types/Diary'
import { client } from './client'

/**
 * diary write.
 *
 * @param params params.
 */
export const diaryWrite = async (params: Diary): Promise<Diary> => {
  const { data } = await client.post<Diary>('/api/diary/write', params)

  return data
}
