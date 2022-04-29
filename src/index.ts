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

export default (target: Target): PromptItem[] => {
  const names = Array.isArray(target) ? target : [target];
  return names.map((name) => prompts[name]);
};
