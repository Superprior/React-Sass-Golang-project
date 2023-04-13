import React, { useEffect, useState } from 'react'
import callAuthAdminAPI from '../lib/callAuthAdminAPI';
import postCodeSample from '../lib/postCodeSample';


interface Props {}

function AdminPage(props: Props): JSX.Element {
    const [access, setAccess] = useState<Boolean>(false);
    const [pwdError, setPwdError] = useState<Boolean>(false);

    useEffect(() => {
        var codeArea = document.getElementById("content") as HTMLInputElement;

        codeArea?.addEventListener('keydown', (e: KeyboardEvent) => {
            if (e.key === "Tab") {
                e.preventDefault();

                if (codeArea.selectionStart) {
                    codeArea.setRangeText(
                        "    ", 
                        codeArea.selectionStart, 
                        codeArea.selectionStart, 
                        'end'
                    );
                };
            };
        });
    }, []);

    const handleSubmitPassword = async (event: any) => {
        event.preventDefault();
     
        const { pwd } = document.forms[0];
        
        const hasAccess = await callAuthAdminAPI({ pwd: pwd.value });
        if (hasAccess === true) {
            setPwdError(false);
            setAccess(true);
        } else {
            setPwdError(true);
        }
    }

    const handleCreateSample = async (event: any) => {
        event.preventDefault();
        const { title_, content } = document.forms[0];

        const created = await postCodeSample({ title: title_.value, langSlug: 'py', content: content.value });
        if (created == true) {
            // handle this...
        } else {
            // and handle errors...
        }
    }

    return (
        <>
            { access ? (
                <div style={{ paddingInline: '20px' }}>
                    <h1>Manage website</h1>
                    <form onSubmit={handleCreateSample}>
                        <div style={{ display: 'block' }}>
                            <p style={{ margin: '0' }}>Title</p>
                            <input name="title_" required />
                        </div>
                        <div style={{ display: 'block' }}>
                            <p style={{ margin: '0' }}>Content (python only)</p>
                            <textarea name="content" id="content" rows={20} cols={100} required></textarea>
                        </div>
                        <div>
                            <input type="submit" />
                        </div>
                    </form>
                </div>
            ) : (
                <>
                    <h1>Admin page</h1>
                    <form onSubmit={handleSubmitPassword}>
                        <input type="password" name="pwd" required style={{ backgroundColor: `${pwdError ? 'red' : 'white '}`}}/>
                        <div>
                            <input type="submit"/>
                        </div>
                    </form>
                </>
            )}
        </>
    )
}

export default AdminPage;