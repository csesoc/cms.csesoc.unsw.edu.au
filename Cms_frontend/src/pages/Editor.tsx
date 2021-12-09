import React, { useState } from 'react';
import styled from 'styled-components';

import EditorHeader from 'src/components/Editor/EditorHeader';
import EditorFile from 'src/components/Editor/EditorFile';
import EditorSidebar from 'src/components/Editor/EditorSidebar';

import { EditorState } from "draft-js";

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
`

const Editor: React.FC = () => {
  const [editorState, setEditorState] = useState(() => {
    return EditorState.createEmpty();
  });

  return (
    <div style={{ height: '100%' }}>
      <EditorHeader />
      <div className="Editor" style={{ display: 'flex' }}>
        {/* <EditorSidebar
          editorState={editorState}
          setEditorState={setEditorState} /> */}
        <div style={{ flex: 1 }}>
          <Container>
            <EditorFile
              editorState={editorState}
              setEditorState={setEditorState} />
          </Container>
        </div>
      </div>
    </div>
  );
};

export default Editor;