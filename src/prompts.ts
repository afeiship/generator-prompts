import yoHelper from '@jswork/yeoman-generator-helper';
import nxNpmRegistries from '@afeiship/next-npm-registries';

const NPM_CHOICES = ['npm', 'github', 'alo7'].map((item) => ({
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

export const component_name = {
  type: 'input',
  name: 'component_name',
  message: 'Your component_name(rc-button)?'
};

export const hook_name = {
  type: 'input',
  name: 'hook_name',
  message: 'Your hook_name(use-button)?'
};

export const description = {
  type: 'input',
  name: 'description',
  message: 'Your description?',
  validate: Boolean
};
