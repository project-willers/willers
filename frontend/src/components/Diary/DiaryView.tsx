import { userAtom } from '@/states/auth'
import { editDiaryAtom } from '@/states/diaries'
import { Diary } from '@/types/Diary'
import { Edit } from '@mui/icons-material'
import {
  Button,
  Card,
  CardContent,
  CardHeader,
  Divider,
  Grid,
} from '@mui/material'
import { useAtom } from 'jotai'
import Link from 'next/link'
import ReactMarkdown from 'react-markdown'

/**
 * DiaryView props.
 */
export type DiaryViewProps = {
  diary: Diary
}

/**
 * DiaryView component.
 */
export const DiaryView: React.VFC<DiaryViewProps> = (props) => {
  const { diary } = props
  const [user] = useAtom(userAtom)
  const [, setEditDiary] = useAtom(editDiaryAtom)
  const dateText = new Date(diary.selectAt).toLocaleDateString()

  return (
    <>
      <Card>
        <CardHeader
          title={dateText}
          subheader={
            <Link href={`/diaries/${diary.name}`} passHref>
              <Button component="a">投稿者: {diary.name}</Button>
            </Link>
          }
        />
        <Divider />
        <CardContent>
          {user.name === diary.name && (
            <Button startIcon={<Edit />} onClick={() => setEditDiary(diary)}>
              編集
            </Button>
          )}
          <ReactMarkdown>{diary.content}</ReactMarkdown>
        </CardContent>
      </Card>
    </>
  )
}
