import * as prompts from './prompts';

type Target = string | string[];
interface PromptItem {
  type: string;
  name: string;
  message: string;
  default?: unknown;
  validate?: () => Boolean;
}

export const supports = (): string[] => Object.keys(prompts);

// https://github.com/puppeteer/puppeteer/issues/6852
// if use export default: error TS2309: An export assignment cannot be used in a module with other exported elements.
module.exports = (target: Target): PromptItem[] => {
  const names = Array.isArray(target) ? target : [target];
  return names.map((name) => prompts[name]);
};
