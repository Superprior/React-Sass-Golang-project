import React, { useEffect, useState } from 'react'
import useAdminAccess from '../lib/callAuthAdminAPI';
import callAuthAdminAPI from '../lib/callAuthAdminAPI';


interface Props {}

function AdminPage(props: Props): JSX.Element {
    const [access, setAccess] = useState<Boolean>(false);
    const [pwdError, setPwdError] = useState<Boolean>(false);

    const handleSubmit = async (event: any) => {
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

    return (
        <>
            { access ? (
                <h1>Manage website</h1>
            ) : (
                <>
                    <h1>Admin page</h1>
                    <form onSubmit={handleSubmit}>
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