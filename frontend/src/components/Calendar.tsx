import { Grid, Paper, Typography } from '@mui/material'

/**
 * Calendar props.
 */
export type CalendarProps = {}

const weeks = ['日', '月', '火', '水', '木', '金', '土']

/**
 * Calendar component.
 */
export const Calendar: React.VFC<CalendarProps> = (props) => {
  return (
    <>
      <Grid container>
        {weeks.map((week) => (
          <Grid key={week} item xs={12 / weeks.length}>
            <Paper variant="outlined" square elevation={0} sx={{ p: 1 }}>
              <Typography textAlign="center">{week}</Typography>
            </Paper>
          </Grid>
        ))}
      </Grid>
    </>
  )
}
