import { DiaryView } from '@/components/Diary/DiaryView'
import { AppLayout } from '@/layouts/AppLayout/AppLayout'
import { diariesAtom } from '@/states/diaries'
import { useAtom } from 'jotai'
import { useRouter } from 'next/router'

/**
 * Diary component.
 */
export const UserDiaryView: React.VFC = (props) => {
  const [diaries] = useAtom(diariesAtom)
  const router = useRouter()

  const { userName, year, month, day } = router.query as {
    userName: string
    year: string
    month: string
    day: string
  }

  const diary = diaries.find((diary) => {
    const select = new Date(diary.selectAt)

    return (
      diary.name === userName &&
      select.getFullYear() === parseInt(year) &&
      select.getMonth() + 1 === parseInt(month) &&
      select.getDate() === parseInt(day)
    )
  })

  return <AppLayout>{diary && <DiaryView diary={diary} />}</AppLayout>
}
