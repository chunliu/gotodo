import * as msal from "@azure/msal-browser";

const msalConfig = {
    auth: {
        clientId: process.env.REACT_APP_AAD_CLIENT_ID as string,
        authority: 'https://login.microsoftonline.com/' + process.env.REACT_APP_AAD_TENANT_ID as string,
    }
};

const aadAuth = new msal.PublicClientApplication(msalConfig);

export default aadAuth;