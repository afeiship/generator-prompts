import * as prompts from './prompts';

type Target = string | string[];
interface PromptItem {
  type: string;
  name: string;
  message: string;
  default?: unknown;
  validate?: () => Boolean;
}

// const supports = (): string[] => Object.keys(prompts);
// 2 者同时使用的时候，不能使用 `export =` 这种用法
// https://github.com/puppeteer/puppeteer/issues/6852
// if use export default: error TS2309: An export assignment cannot be used in a module with other exported elements.
export default (target: Target): PromptItem[] => {
  const names = Array.isArray(target) ? target : [target];
  return names.map((name) => prompts[name]);
};
