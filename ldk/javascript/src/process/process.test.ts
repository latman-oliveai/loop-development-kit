import { mocked } from 'ts-jest/utils';
import * as process from '.';

describe('Process', () => {
  beforeEach(() => {
    oliveHelps.process = {
      all: jest.fn(),
      listenAll: jest.fn(),
    };
  });

  describe('all', () => {
    it('returns a promise result with expected proces infos', () => {
      const expected: process.ProcessInfo[] = [
        {
          arguments: 'arguments',
          command: 'command',
          pid: 456,
        },
      ];
      mocked(oliveHelps.process.all).mockImplementation((callback) =>
        callback(undefined, expected),
      );

      const actual = process.all();

      return expect(actual).resolves.toBe(expected);
    });

    it('returns a rejected promise', () => {
      const exception = 'Exception';
      mocked(oliveHelps.process.all).mockImplementation(() => {
        throw exception;
      });

      const actual = process.all();

      return expect(actual).rejects.toBe(exception);
    });
  });

  describe('listenAll', () => {
    it('passed in listen all callback to olive helps', () => {
      const callback = jest.fn();
      process.listenAll(callback);

      expect(oliveHelps.process.listenAll).toHaveBeenCalledWith(callback, expect.any(Function));
    });

    it('rejects with the error when the underlying call throws an error', () => {
      const exception = 'Exception';
      mocked(oliveHelps.process.listenAll).mockImplementation(() => {
        throw exception;
      });

      const callback = jest.fn();
      expect(() => process.listenAll(callback)).rejects.toBe(exception);
    });
  });
});
