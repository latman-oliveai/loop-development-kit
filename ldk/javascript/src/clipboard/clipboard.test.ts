import { mocked } from 'ts-jest/utils';
import * as clipboard from '.';

describe('Clipboard', () => {
  beforeEach(() => {
    oliveHelps.clipboard = {
      includeOliveHelpsEvents: jest.fn(),
      read: jest.fn(),
      write: jest.fn(),
      listen: jest.fn(),
    };
  });

  describe('read', () => {
    it('returns a promise result with expected clipboard value', () => {
      const expected = 'expected string';
      mocked(oliveHelps.clipboard.read).mockImplementation((callback) => callback(undefined, expected));

      const actual = clipboard.read();

      return expect(actual).resolves.toBe(expected);
    });

    it('returns a rejected promise', () => {
      const exception = 'Exception';
      mocked(oliveHelps.clipboard.read).mockImplementation(() => {
        throw exception;
      });

      const actual = clipboard.read();

      return expect(actual).rejects.toBe(exception);
    });
  });

  describe('listen', () => {
    it('sets clipboard olive helps configuration', () => {
      const includeOliveHelpsEvents = true;
      const callback = jest.fn();
      clipboard.listen(includeOliveHelpsEvents, callback);

      expect(oliveHelps.clipboard.includeOliveHelpsEvents).toHaveBeenCalledWith(
        includeOliveHelpsEvents,
      );
    });

    it('passed in listen function to olive helps', () => {
      const callback = jest.fn();
      clipboard.listen(true, callback);

      expect(oliveHelps.clipboard.listen).toHaveBeenCalledWith(callback, expect.any(Function));
    });

    it('rejects with the error when the underlying call throws an error', () => {
      const exception = 'Exception';
      mocked(oliveHelps.clipboard.listen).mockImplementation(() => {
        throw exception;
      });

      const callback = jest.fn();
      expect(() => clipboard.listen(false, callback)).rejects.toBe(exception);
    });
  });

  describe('write', () => {
    it('writes text to an olive helps clipboard', () => {
      const expectedText = 'text';
      mocked(oliveHelps.clipboard.write).mockImplementation((text, callback) => {
        callback(undefined);
      });

      const actual = clipboard.write(expectedText);

      expect(oliveHelps.clipboard.write).toHaveBeenCalledWith(expectedText, expect.any(Function));
      return expect(actual).resolves.toBeUndefined();
    });

    it('returns a rejected promise', () => {
      const exception = 'Exception';
      mocked(oliveHelps.clipboard.write).mockImplementation(() => {
        throw exception;
      });

      const actual = clipboard.write('text');

      return expect(actual).rejects.toBe(exception);
    });
  });
});
