// @ts-ignore
import fn, { supports } from '../src';

describe('api.basic', () => {
  test('generate 2 props', () => {
    const res = fn(['project_name', 'description']);
    expect(res.length).toBe(2);
  });
});
