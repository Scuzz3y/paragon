import { useQuery } from "@apollo/react-hooks";
import gql from "graphql-tag";
import * as React from "react";
import { Feed, Loader } from "semantic-ui-react";
import { EventKind, XEvent, XNoEventsFound } from "../components/event";
import { XErrorMessage } from "../components/messages";
import { Event } from "../graphql/models";

export const EVENT_FEED_QUERY = gql`
  {
    events(input: { limit: 10 }) {
      id
      creationTime
      kind

      job {
        id
        name
      }
      file {
        id
        name
        size
      }
      credential {
        id
        secret
        principal
        kind
      }
      link {
        id
        alias
        expirationTime
        clicks
      }
      tag {
        id
        name
      }
      target {
        id
        name
      }
      task {
        id
        claimTime
        execStartTime
        execStopTime
      }
      user {
        id
        name
        isActivated
        isAdmin
      }
      service {
        id
        name
        isActivated
      }
      event {
        id
        kind
      }
      likers {
        id
      }
      owner {
        id
        name
        isActivated
        isAdmin
      }
      svcOwner {
        id
        name
        isActivated
      }
    }
  }
`;

export type EvnetFeedResponse = {
  events: Event[];
};

const XEventFeedView = () => {
  const { called, loading, error, data } = useQuery<EvnetFeedResponse>(
    EVENT_FEED_QUERY,
    {
      pollInterval: 3000
    }
  );

  if (!data || !data.events || data.events.length < 1) {
    return <XNoEventsFound />;
  }

  return (
    <React.Fragment>
      <Loader disabled={!called || !loading} />
      <Feed size="large">
        {data.events.map((e, index) => {
          console.log(e.kind);
          console.log(EventKind[e.kind]);
          return <XEvent key={index} event={e} kind={EventKind[e.kind]} />;
        })}
      </Feed>
      <XErrorMessage title="Error Loading Feed" err={error} />
    </React.Fragment>
  );
};

export default XEventFeedView;
