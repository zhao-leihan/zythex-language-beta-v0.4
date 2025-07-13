import * as vscode from 'vscode';
import * as path from 'path';

export function activate(context: vscode.ExtensionContext) {
  console.log('Zythex extension is now active!');

  const runZythxCommand = vscode.commands.registerCommand('zythex.runFile', () => {
    const editor = vscode.window.activeTextEditor;
    if (!editor) {
      vscode.window.showErrorMessage('No active editor');
      return;
    }

    const filePath = editor.document.fileName;
    const workspaceFolder = vscode.workspace.workspaceFolders?.[0].uri.fsPath || '';
    const exePath = path.join(workspaceFolder, 'zythx.exe');

    const terminal = vscode.window.createTerminal('Zythex Runner');
    terminal.show();
    terminal.sendText(`& "${exePath}" run "${filePath}"`);

  });

  context.subscriptions.push(runZythxCommand);
}

export function deactivate() {}
