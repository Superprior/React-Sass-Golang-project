import React from 'react'


interface Props {
    pwd: string;
}

const callAuthAdminAPI = async (props: Props): Promise<boolean> => {
    try {
        const url = `${process.env.REACT_APP_PROTOCOL}://${process.env.REACT_APP_HOSTNAME}:8080/api/auth_admin`;
        const responseData = await fetch(url, {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({"Pwd": props.pwd})
        });
        if (responseData.status === 200) {
            return true;
        } else if (responseData.status === 400) {
            /**
             * @todo Handle empty password error
             */

            return false;
        } else if (responseData.status === 401) {
            /**
             * @todo Handle wrong password error
             */
            return false;
        } else {
            return false;
        }
    } catch (error) {
        return false;
    }
}

export default callAuthAdminAPI;