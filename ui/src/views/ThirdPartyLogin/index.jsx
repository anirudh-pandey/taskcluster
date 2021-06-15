import React, { Component } from 'react';
import { graphql } from 'react-apollo';
import { scopeIntersection } from 'taskcluster-lib-scopes';
import { parse } from 'qs';
import Typography from '@material-ui/core/Typography';
import Spinner from '../../components/Spinner';
import AuthConsent from '../../components/AuthConsent';
import Dashboard from '../../components/Dashboard';
import { withAuth } from '../../utils/Auth';
import fromNowJSON from '../../utils/fromNowJSON';
import thirdPartyLoginQuery from './thirdPartyLogin.graphql';

@withAuth
@graphql(thirdPartyLoginQuery, {
  skip: ({ user }) => !user,
  options: () => ({
    fetchPolicy: 'network-only',
  }),
})
export default class ThirdPartyLogin extends Component {
  query = this.props.location.search.slice(1);

  parsedQuery = parse(this.query);

  state = {
    formData: null,
  };

  static getDerivedStateFromProps(props, state) {
    const {
      data,
      user,
      location: { search },
    } = props;
    const query = parse(search.slice(1));
    const scopes = (query.scope || []).split(' ');
    const registeredClientId = query.client_id;

    if (
      !data ||
      state.formData ||
      !(data.currentScopes instanceof Array) ||
      !query.transactionID
    ) {
      return null;
    }

    return {
      formData: {
        ...state.formData,
        scopes: scopeIntersection(scopes, data.currentScopes),
        expires: fromNowJSON(query.expires),
        description: `Client generated by ${user.credentials.clientId} for OAuth2 Client ${registeredClientId}`,
      },
    };
  }

  handleExpirationChange = expires => {
    this.setState({
      formData: {
        ...this.state.formData,
        expires,
      },
    });
  };

  handleInputChange = ({ target: { name, value } }) => {
    this.setState({
      formData: {
        ...this.state.formData,
        [name]: value,
      },
    });
  };

  handleScopesChange = scopes => {
    this.setState({
      formData: {
        ...this.state.formData,
        scopes,
      },
    });
  };

  render() {
    const { user, data } = this.props;
    const { formData } = this.state;
    const { origin } = new URL(
      decodeURIComponent(this.parsedQuery.redirect_uri)
    );

    return (
      <Dashboard title="Third Party Login">
        {data && data.loading && <Spinner loading />}
        {formData && (
          <AuthConsent
            transactionID={this.parsedQuery.transactionID}
            registeredClientId={this.parsedQuery.client_id}
            clientId={this.parsedQuery.clientId}
            onExpirationChange={this.handleExpirationChange}
            onInputChange={this.handleInputChange}
            onScopesChange={this.handleScopesChange}
            formData={formData}
          />
        )}
        {!user && (
          <Typography variant="subtitle1">
            Sign in to provide credentials to <code>{origin}</code>
          </Typography>
        )}
      </Dashboard>
    );
  }
}
