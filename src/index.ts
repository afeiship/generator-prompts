import * as prompts from './prompts';

module.exports = (inNames) => {
  return inNames.map((name) => prompts[name]);
};
