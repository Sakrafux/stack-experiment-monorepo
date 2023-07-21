import { updateCurrentUser } from 'api';
import { useLoginContext } from 'context';
import { useLogout } from 'hooks';
import { UpdateUserDto } from 'models';
import { useEffect, useState } from 'react';

const Settings = () => {
  const [value, setValue] = useState<UpdateUserDto>({ username: '', email: '', bio: '', image: '', password: '' });

  const { user } = useLoginContext().state;

  const logout = useLogout();

  useEffect(() => {
    if (user) {
      const { username, email, bio, image } = user;
      setValue({ username, email, bio, image, password: '' });
    }
  }, [user]);

  return (
    <div className="settings-page">
      <div className="container page">
        <div className="row">
          <div className="col-md-6 offset-md-3 col-xs-12">
            <h1 className="text-xs-center">Your Settings</h1>

            <form onSubmit={e => e.preventDefault()}>
              <fieldset>
                <fieldset className="form-group">
                  <input
                    className="form-control"
                    type="text"
                    placeholder="URL of profile picture"
                    value={value.image}
                    onChange={e => setValue({ ...value, image: e.target.value })}
                  />
                </fieldset>
                <fieldset className="form-group">
                  <input
                    className="form-control form-control-lg"
                    type="text"
                    placeholder="Your Name"
                    value={value.username}
                    onChange={e => setValue({ ...value, username: e.target.value })}
                  />
                </fieldset>
                <fieldset className="form-group">
                  <textarea
                    className="form-control form-control-lg"
                    rows={8}
                    placeholder="Short bio about you"
                    value={value.bio ?? ''}
                    onChange={e => setValue({ ...value, bio: e.target.value })}
                  ></textarea>
                </fieldset>
                <fieldset className="form-group">
                  <input
                    className="form-control form-control-lg"
                    type="text"
                    placeholder="Email"
                    value={value.email}
                    onChange={e => setValue({ ...value, email: e.target.value })}
                  />
                </fieldset>
                <fieldset className="form-group">
                  <input
                    className="form-control form-control-lg"
                    type="password"
                    placeholder="Password"
                    value={value.password}
                    onChange={e => setValue({ ...value, password: e.target.value })}
                  />
                </fieldset>
                <button className="btn btn-lg btn-primary pull-xs-right" onClick={() => updateCurrentUser(value)}>
                  Update Settings
                </button>
              </fieldset>
            </form>
            <hr />
            <button className="btn btn-outline-danger" onClick={() => logout()}>
              Or click here to logout.
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Settings;
