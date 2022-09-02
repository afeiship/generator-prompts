import yoHelper from '@jswork/yeoman-generator-helper';
import nxNpmRegistries from '@jswork/next-npm-registries';
import '@jswork/next-git-url';
import gitInfo from '@jswork/git-info';
const gitUrl = new nx.GitUrl(gitInfo.url());
const SPT_REGIESTRY = ['npm', 'github', 'alo7', 'null'];

// special for next
declare var nx: any;

// base
const NPM_CHOICES = SPT_REGIESTRY.map((item) => ({
  name: item,
  value: nxNpmRegistries(item)
}));

export const scope = {
  type: 'input',
  name: 'scope',
  message: 'Your project_name scope (eg: `@babel`)?',
  default: 'jswork'
};

export const registry = {
  type: 'list',
  name: 'registry',
  message: 'Your registry',
  choices: NPM_CHOICES
};

export const project_name = {
  type: 'input',
  name: 'project_name',
  message: 'Your project_name (eg: like this `my-project` )?',
  default: yoHelper.discoverRoot
};

export const project_url = {
  type: 'input',
  name: 'project_url',
  message: 'Your project_url?',
  default: gitUrl.url
};

export const description = {
  type: 'input',
  name: 'description',
  message: 'Your description?',
  validate: Boolean
};

export const homepage = {
  type: 'input',
  name: 'homepage',
  message: 'Your homepage?',
  default: 'https://js.work'
};

export const author = {
  type: 'input',
  name: 'author',
  message: 'Your author?',
  default: gitInfo.username()
};

export const email = {
  type: 'input',
  name: 'email',
  message: 'Your email?',
  default: gitInfo.email()
};

export const filetype = {
  type: 'list',
  name: 'filetype',
  message: 'Your filetype?',
  choices: [
    { name: 'TypeScript', value: 'ts' },
    { name: 'JavaScript', value: 'js' }
  ]
};

// ----- react
export const component_name = {
  type: 'input',
  name: 'component_name',
  message: 'Your component_name(rc-button)?'
};

export const module_name = {
  type: 'input',
  name: 'module_name',
  message: 'Your module_name(mod1)?'
};

export const resource_name = {
  type: 'input',
  name: 'resource_name',
  message: 'Your resource_name(users)?'
};

export const hook_name = {
  type: 'input',
  name: 'hook_name',
  message: 'Your hook_name(use-button)?'
};
