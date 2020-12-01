
export interface FileDependency {
  type: 'core/file'
  path: string
}

export interface AptDependency {
  type: 'core/apt'
  path: string
}

export interface FolderDependency {
  type: 'core/folder'
  path: string
}

export interface EnvVarDependency {
  type: 'core/envvar',
  name: string,
  value?: string
}

export interface CarpGroupDependency {
  type: 'core/carpgroup',
  name: string
}

export type Dependency =
  FileDependency |
  AptDependency |
  FolderDependency |
  EnvVarDependency |
  CarpGroupDependency

export interface GroupDefinition {
  requires: Dependency[]
}

export interface CarpFile {
  groups: {
    [key: string]: GroupDefinition
  }
}

const isDependency = (val: any): val is Dependency => {
  return true // check
}

const isGroup = (val: any): val is GroupDefinition => {
  if (!val.requires || !Array.isArray(val.requires)) {
    return false
  }

  for (const inner of val.requires) {
    if (!isDependency(inner)) {
      return false
    }
  }

  return true
}

export const isCarpfile = (val: any): val is CarpFile => {
  if (!('groups' in val)) {
    return false
  }

  if (val.groups === null || typeof val.groups !== 'object') {
    return false
  }

  for (const inner of Object.values(val.groups)) {
    if (!isGroup(inner)) {
      return false
    }
  }

  return true
}
