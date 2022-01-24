import * as prompts from './prompts';

export default (inNames): void => {
  return inNames.map((name) => prompts[name]);
};
