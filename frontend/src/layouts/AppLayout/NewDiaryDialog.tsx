import { DiaryEditDialog } from '@/components/Diary/DiaryEditDialog/DiaryEditDialog'

/**
 * NewDiaryDialog props.
 */
export type NewDiaryDialogProps = {
  open: boolean
  onClose: () => void
}

/**
 * NewDiaryDialog component.
 */
export const NewDiaryDialog: React.VFC<NewDiaryDialogProps> = (props) => {
  return <DiaryEditDialog open={props.open} onClose={props.onClose} />
}
