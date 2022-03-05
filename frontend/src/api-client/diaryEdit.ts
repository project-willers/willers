import { Diary } from '@/types/Diary'
import { client } from './client'

/**
 * diary edit.
 *
 * @param params params.
 */
export const diaryEdit = async (params: Diary): Promise<Diary> => {
  const { data } = await client.post<Diary>('/api/diary/edit', params)

  return data
}
