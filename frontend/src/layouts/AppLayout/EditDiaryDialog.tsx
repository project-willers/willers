import { DiaryEditDialog } from '@/components/Diary/DiaryEditDialog/DiaryEditDialog'
import { Diary } from '@/types/Diary'

/**
 * EditDiaryDialog props.
 */
export type EditDiaryDialogProps = {
  open: boolean
  onClose: () => void
  diary: Diary
  onSave?(): void
}

/**
 * EditDiaryDialog component.
 */
export const EditDiaryDialog: React.VFC<EditDiaryDialogProps> = (props) => {
  return (
    <DiaryEditDialog
      open={props.open}
      onClose={props.onClose}
      onSave={props.onSave}
      diary={props.diary}
    />
  )
}
