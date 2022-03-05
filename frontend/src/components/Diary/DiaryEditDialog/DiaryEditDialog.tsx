import { diaryEdit } from '@/api-client/diaryEdit'
import { diaryWrite } from '@/api-client/diaryWrite'
import { userAtom } from '@/states/auth'
import { Diary } from '@/types/Diary'
import { Send } from '@mui/icons-material'
import { DatePicker } from '@mui/lab'
import {
  Button,
  Dialog,
  DialogContent,
  DialogTitle,
  FormControl,
  Grid,
  InputLabel,
  MenuItem,
  Select,
  TextField,
} from '@mui/material'
import { Box } from '@mui/system'
import { format } from 'date-fns'
import { useAtom } from 'jotai'
import { useEffect, useState } from 'react'

/**
 * DiaryEditDialog props.
 */
export type DiaryEditDialogProps = {
  open: boolean
  onClose: () => void
  onSave?(): void
  diary?: Diary
}

type Template = {
  name: string
  content: string
}

const templates: Template[] = [
  {
    name: '無地',
    content: '',
  },
]

/**
 * DiaryEditDialog component.
 */
export const DiaryEditDialog: React.VFC<DiaryEditDialogProps> = (props) => {
  const { diary } = props
  const [date, setDate] = useState<Date | null>(
    diary?.selectAt ? new Date(diary.selectAt) : new Date()
  )
  const [template, setTemplate] = useState<Template>(templates[0])
  const [content, setContent] = useState(diary?.content ?? '')
  const [user] = useAtom(userAtom)

  useEffect(() => {
    if (diary) {
      setDate(new Date(diary.selectAt))
      setTemplate(templates.find((t) => t.name === diary.name) ?? templates[0])
      setContent(diary.content)
    }
  }, [diary])

  const onSendClicked = async () => {
    if (!date) {
      return
    }

    if (!user) {
      return
    }

    await diaryEdit({
      name: user.name,
      content,
      selectAt: format(date, 'yyyy-MM-dd HH:mm:ss'),
      updatedAt: format(new Date(), 'yyyy-MM-dd HH:mm:ss'),
    })

    if (props.onSave) {
      props.onSave()
    }
  }

  return (
    <Dialog maxWidth="md" fullWidth open={props.open} onClose={props.onClose}>
      <DialogContent>
        <Grid container spacing={2}>
          <Grid item xs={6}>
            <DatePicker
              label="日時"
              value={date}
              onChange={(v) => setDate(v)}
              renderInput={(props) => <TextField {...props} fullWidth />}
            />
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth>
              <InputLabel id="diary-edit-dialog-template">
                テンプレート
              </InputLabel>
              <Select
                labelId="diary-edit-dialog-template"
                label="テンプレート"
                value={template.name}
                onChange={(t) =>
                  setTemplate(
                    templates.find(({ name }) => name === t.target.value) ??
                      templates[0]
                  )
                }
              >
                {templates.map((t) => (
                  <MenuItem key={t.name} value={t.name}>
                    {t.name}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <TextField
              label="内容"
              multiline
              rows={20}
              fullWidth
              value={content}
              onChange={(e) => setContent(e.target.value)}
            />
          </Grid>
          <Grid item xs={12}>
            <Box sx={{ display: 'flex' }}>
              <Box flex={1} />
              <Button
                variant="contained"
                disableElevation
                startIcon={<Send />}
                onClick={onSendClicked}
              >
                送信
              </Button>
            </Box>
          </Grid>
        </Grid>
      </DialogContent>
    </Dialog>
  )
}
