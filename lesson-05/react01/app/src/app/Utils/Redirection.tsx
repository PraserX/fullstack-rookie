import * as React from 'react';
import { Redirect } from 'react-router';

const RenderRedirect = (execute, path) => execute ? <Redirect to={path} /> : null
