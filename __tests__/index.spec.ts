import fn from '../src';

describe('api.basic', () => {
  test('normail single value case', () => {
    const res = fn(['project_name', 'description']);
    console.log(res);
  });
});
