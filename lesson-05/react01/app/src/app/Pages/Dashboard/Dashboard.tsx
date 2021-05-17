import * as React from 'react';
import {
  Breadcrumb,
  BreadcrumbItem,
  PageSection,
  PageSectionVariants,
  Title,
  Text,
  TextInput,
  TextArea,
  TextContent,
  TextVariants,
  Spinner,
  EmptyState,
  EmptyStateBody,
  Card,
  CardTitle,
  CardBody,
  CardFooter,
  Grid,
  GridItem,
  Form,
  FormGroup,
  ActionGroup,
  Button
} from '@patternfly/react-core';
import { Redirect } from 'react-router';

import { HandleError } from '@app/Handlers/ErrorHandler';

import { ICommentResponse } from '@app/Types/Responses/Comment';
import { INotificationContent, FCNotifications } from '@app/Components/Notifications';
import { ErrorUnwantedRedirect, ErrorNetwork, ErrorUnexpected, IResponse, APIGetComments, APIPostRequest, ErrorForbidden, ErrorNotFound } from '@app/Handlers/ApiHandler';
import { GetReadableDate, GetReadableTime } from '@app/Utils/Helpers';
import { WindowsIcon } from '@patternfly/react-icons';

interface IStats {
  usersActive: number;
  usersLogged: number;
  usersSum: number;
  usersOnAccount: number;
  usersOnPortal: number;
  usersOnNextcloud: number;
  usersOnMisp: number;
  usersOnIam: number;
}

const Dashboard: React.FunctionComponent = () => {
  const [execRedirect, setExecRedirect] = React.useState(false)
  const [destPath, setDestPath] = React.useState("")

  const [loadedComments, setLoadedComments] = React.useState(false)
  const [loading, setLoading] = React.useState(true)

  const [comments, setComments] = React.useState(Array<ICommentResponse>())
  const [refresh, setRefresh] = React.useState(false);
  const [update, setUpdate] = React.useState(false);

  const [notifications, setNotifications] = React.useState(Array<INotificationContent>())

  const [formNickname, setFormNickname] = React.useState("")
  const [formEmail, setFormEmail] = React.useState("")
  const [formComment, setFormComment] = React.useState("")

  React.useEffect(() => {
    var nickname = window.localStorage.getItem('nickname')
    var email = window.localStorage.getItem('email')
    
    setFormNickname(nickname !== null ? nickname : "")
    setFormEmail(email !== null ? email : "")
  }, []);

  // React.useEffect(() => {
  //   if (!refresh) {
  //     setInterval(() => { setRefresh(!refresh) }, 5000);
  //   }
  // }, [refresh]);

  React.useEffect(() => {
    APIGetComments(setComments, comments)
      .then(status => status.ok && setLoadedComments(true))
      .catch(error => {
        setComments(Array<ICommentResponse>())
        handleError(error)
      })
  }, [refresh, update]);

  // Finish loading
  React.useEffect(() => {
    if (!loadedComments) {
      return
    }

    setLoading(false)
  }, [loadedComments])

  const handleError = (error: IResponse) => {
    if (error.error === ErrorUnwantedRedirect) {
      console.info("Unwanted redirect detected")
    } else if ((error.error === ErrorForbidden) || (error.error === ErrorNotFound) || (error.error === ErrorUnexpected)) {
      HandleError(false, error, handleRedirect)
    } else if (error.error === ErrorNetwork) {
      HandleError(true, null, handleRedirect)
    } else {
      console.error("Something unexpected happened")
    }
  }

  /**
   * 
   * @param path 
   */
  const handleRedirect = (path: string): void => {
    setDestPath(path)
    setExecRedirect(true)
  }

  const handleSaveAction = () => {
    window.localStorage.setItem('nickname', formNickname);
    window.localStorage.setItem('email', formEmail);

    var notification: INotificationContent = { variant: "success", title: "User information update success" }
    setNotifications(Array(...notifications, notification))
  }

  const getFTime = (timestamp: string) => {
    var ts = new Date(timestamp)
    return (
      <Text component={TextVariants.small}>
        Dne {GetReadableDate(ts) + " v " + GetReadableTime(ts)} hodin
      </Text>
    )
  }

  const getComment = (item: ICommentResponse) => {
    return (
      <Card>
        <CardTitle>{item.user.nickname}, {item.user.email}</CardTitle>
        <CardBody>{item.comment}</CardBody>
        <CardFooter>{getFTime(item.timestamp)}</CardFooter>
      </Card>
    )
  }

  const handleSubmitComment = () => {
    var newComment = {
      "comment": formComment,
      "nickname": window.localStorage.getItem('nickname'),
      "email": window.localStorage.getItem('email')
    }

    APIPostRequest("/api/v1/comments", "", JSON.stringify(newComment))
      .then(response => {
        if (response.ok) {
          var notification: INotificationContent = { variant: "success", title: "Comment has been sent" }
          setNotifications(Array(...notifications, notification))
          setFormComment("")
          setUpdate(!update)
        }
      }).catch(error => { 
        console.log(error)
        var notification: INotificationContent = { variant: "success", title: "Unexpected error! Comment has not been sent" }
        setNotifications(Array(...notifications, notification))
      })
  }

  const redirection = execRedirect
    ? <Redirect to={destPath} />
    : <React.Fragment></React.Fragment>

  return (
    <React.Fragment>
      {redirection}
      <FCNotifications notifications={notifications} setNotificationsCallback={setNotifications}></FCNotifications>
      <PageSection variant={PageSectionVariants.darker} >
        <Title headingLevel="h1" size="xl">Comments</Title>
      </PageSection>
      <PageSection variant={PageSectionVariants.light}>
        {loading &&
          <EmptyState>
            <Title headingLevel="h5" size="lg">
              Loading ...
          </Title>
            <EmptyStateBody>
              <Spinner size="lg" />
            </EmptyStateBody>
          </EmptyState>
        }
        {!loading &&
          <React.Fragment>
            <Grid>
              <GridItem sm={12} md={5} lg={4} xl={3} xl2={3} style={{ padding: "0px 10px" }}>
                <Form>
                  <FormGroup label="Name" isRequired fieldId="horizontal-form-name" helperText="Please provide your full name">
                    <TextInput
                      value={formNickname}
                      isRequired
                      type="text"
                      id="horizontal-form-name"
                      aria-describedby="horizontal-form-name-helper"
                      name="horizontal-form-name"
                      onChange={(value) => {setFormNickname(value)}}
                    />
                  </FormGroup>
                  <FormGroup label="Email" isRequired fieldId="horizontal-form-email">
                    <TextInput
                      value={formEmail}
                      onChange={(value) => {setFormEmail(value)}}
                      isRequired
                      type="email"
                      id="horizontal-form-email"
                      name="horizontal-form-email"
                    />
                  </FormGroup>
                  <ActionGroup>
                    <Button variant="primary" onClick={() => handleSaveAction()}>Save</Button>
                  </ActionGroup>
                </Form>
              </GridItem>
              <GridItem sm={12} md={7} lg={8} xl={9} xl2={9} style={{ padding: "0px 10px" }}>
                <Form style={{ padding: "0px 0px 25px 0px" }}>
                  <FormGroup label="Your comment" fieldId="horizontal-form-exp">
                    <TextArea
                      value={formComment}
                      onChange={(value) => {setFormComment(value)}}
                      name="horizontal-form-exp"
                      id="horizontal-form-exp"
                    />
                  </FormGroup>
                  <ActionGroup>
                    <Button variant="primary" onClick={() => handleSubmitComment()}>Send comment</Button>
                    <Button variant="link" onClick={() => setFormComment("")}>Clear</Button>
                  </ActionGroup>
                </Form>
                <React.Fragment>
                  {comments.map(item => getComment(item))}
                </React.Fragment>
              </GridItem>
            </Grid>
          </React.Fragment>
        }
      </PageSection>
    </React.Fragment >
  );
}

export { Dashboard };
