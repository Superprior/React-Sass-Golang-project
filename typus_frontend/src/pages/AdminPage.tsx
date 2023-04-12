import React, { useEffect, useState } from 'react'
import useAdminAccess from '../lib/callAuthAdminAPI';
import callAuthAdminAPI from '../lib/callAuthAdminAPI';
import postCodeSample from '../lib/postCodeSample';


interface Props {}

function AdminPage(props: Props): JSX.Element {
    const [access, setAccess] = useState<Boolean>(false);
    const [pwdError, setPwdError] = useState<Boolean>(false);

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
        const { title_, langSlug, content } = document.forms[0];

        const created = await postCodeSample({ title: title_.value, langSlug: langSlug.value, content: content.value });
        if (created == true) {
            alert('Created!');
        } else {
            alert('Error!');
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
                            <p style={{ margin: '0' }}>Language slug</p>
                            <input name="langSlug" required />
                        </div>
                        <div style={{ display: 'block' }}>
                            <p style={{ margin: '0' }}>Content</p>
                            <textarea name="content" rows={20} cols={100} required></textarea>
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