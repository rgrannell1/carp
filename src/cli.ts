
import docopt from 'docopt'
import carp from './carp.js'

const docs = `
Name:
  carp -

Usage:
  carp --file <path>
`

const cli = async () => {
  const args = docopt.docopt(docs, {})
  await carp(args)
}

cli()
