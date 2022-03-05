import { Diary } from '@/types/Diary'
import { client } from './client'

/**
 * diary read.
 */
export const diaryRead = async (): Promise<Diary[]> => {
  const { data } = await client.get<{
    Diaries: Diary[]
  }>('/api/diary/read')

  return data.Diaries
}
