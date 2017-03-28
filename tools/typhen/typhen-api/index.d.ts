declare module TyphenApi {
  interface integer {}

  type RTMMessage<T> = {
    [P in keyof T]: T[P];
  };
}
