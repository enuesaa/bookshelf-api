import { AiOutlineSwapRight } from 'react-icons/ai'
import Link from 'next/link'
import { useStyles } from '@/styles/use'
import { PageTitle } from '@/components/common/PageTitle'

type ItemProps = {
  title: string,
  id: string,
}
export const Item = ({ title, id }: ItemProps) => {
  const styles = useStyles(theme => ({
    link: theme({ size:'x1', around: 'x2', decorate: 'card' }),
  }))

  return (
    <Link href={`/boards/${id}`} css={styles.link}>
      {title}
    </Link>
  )
}

export const Dashboard = () => {
  const styles = useStyles(theme => ({
    main: theme({ around: 'x1tb' }),
  }))

  return (
    <section css={styles.main}>
      <PageTitle title='Boards'>
        <Link href='/boards'><AiOutlineSwapRight /></Link>
      </PageTitle>
      <ul>
        <Item title='aa' id='aa' />
      </ul>
    </section>
  )
} 