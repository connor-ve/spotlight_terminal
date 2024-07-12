import 'electron';

declare module 'electron' {
  interface IpcRenderer {
    closeCurrentWindow: () => void;
  }
}
