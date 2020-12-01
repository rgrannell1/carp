
import * as fs from 'fs'
import execa from 'execa'
import {
  CarpFile,
  isCarpfile
} from './types.js'

export const readCarpfile = async (fpath: string):Promise<CarpFile> => {
  let isExecutable = false

  try {
    await fs.promises.access(fpath, fs.constants.X_OK)
    isExecutable = true
  } catch (err) { }

  let content
  if (isExecutable) {
    try {
      const { stdout } = await execa(fpath, [])
      content = stdout
    } catch (err) {
      console.log(err)
      process.exit(1)
    }
  } else {
    content = (await fs.promises.readFile(fpath)).toString()
  }

  const parsed = JSON.parse(content)

  if (!isCarpfile(parsed)) {
    throw new Error('failed to parse')
  }

  return parsed
}

interface RawCarpArgs {
  '<path>': string
}

const carp = async (args: RawCarpArgs) => {
  const carpfile = await readCarpfile(args['<path>'])

  console.log(carpfile)

}

export default carp
