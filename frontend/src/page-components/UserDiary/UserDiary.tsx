import { AppLayout } from '@/layouts/AppLayout/AppLayout'
import { jwtAtom } from '@/states/auth'
import { Diary } from '@/types/Diary'
import { diariesAtom } from '@/states/diaries'
import {
  Box,
  Card,
  CardActionArea,
  CardHeader,
  Dialog,
  Grid,
  Paper,
  Typography,
} from '@mui/material'
import { addMonths } from 'date-fns'
import { useAtom } from 'jotai'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import Link from 'next/link'

type Month = {
  year: number
  month: number
  diaries: Diary[]
}

/**
 * UserDiary props.
 */
export type UserDiaryProps = {}

/**
 * UserDiary component.
 */
export const UserDiary: React.VFC<UserDiaryProps> = (props) => {
  const router = useRouter()
  const userName = router.query.userName
  const [date, setDate] = useState(new Date())
  const [diaries] = useAtom(diariesAtom)

  const displayedDiaries =
    diaries?.filter((diary) => {
      const min = addMonths(date, -2).getTime()
      const max = addMonths(date, 2).getTime()
      const select = new Date(diary.selectAt)

      return select.getTime() >= min && select.getTime() <= max
    }) ?? []

  const months: Month[] = []

  for (const diary of displayedDiaries) {
    const date = new Date(diary.selectAt)
    const year = date.getFullYear()
    const month = date.getMonth()
    const monthItem = months.find((m) => m.year === year && m.month === month)

    if (monthItem) {
      monthItem.diaries.push(diary)
    } else {
      months.push({
        year,
        month,
        diaries: [diary],
      })
    }
  }

  return (
    <AppLayout title={`${userName}の日記`}>
      <Grid container spacing={10}>
        {months.map((month, key) => (
          <Grid key={key} item xs={12} md={2.4}>
            <Typography variant="h6" textAlign="center">
              {month.year}
            </Typography>
            <Typography variant="h4" textAlign="center">
              {month.month + 1}月
            </Typography>
            <Box my={1} />
            <Box
              display="flex"
              width="100%"
              alignItems="center"
              flexDirection="column"
            >
              {month.diaries.map((diary, key) => {
                const select = new Date(diary.selectAt)

                return (
                  <Card
                    key={key}
                    sx={{
                      mb: 2,
                      width: '100%',
                      maxWidth: 200,
                      aspectRatio: '1 / 1',
                    }}
                  >
                    <Link
                      href={`/diaries/${userName}/${select.getFullYear()}/${
                        select.getMonth() + 1
                      }/${select.getDate()}`}
                      passHref
                    >
                      <CardActionArea
                        component="a"
                        sx={{
                          width: '100%',
                          height: '100%',
                          display: 'flex',
                          justifyContent: 'center',
                          alignItems: 'center',
                        }}
                      >
                        <Typography variant="h3">
                          {new Date(diary.selectAt).getDate()}日
                        </Typography>
                      </CardActionArea>
                    </Link>
                  </Card>
                )
              })}
            </Box>
          </Grid>
        ))}
      </Grid>
    </AppLayout>
  )
}
