import React from 'react'
import p from 'prop-types'
import { urlParamSelector } from '../selectors'
import { setURLParam } from '../actions'
import gql from 'graphql-tag'
import { FormControlLabel, Switch } from '@material-ui/core'
import { useQuery } from 'react-apollo'
import { useSelector, useDispatch } from 'react-redux'

const tzQuery = gql`
  query($id: ID!) {
    schedule(id: $id) {
      id
      timeZone
    }
  }
`

export function ScheduleTZFilter(props) {
  const params = useSelector(urlParamSelector)
  const zone = params('tz', 'local')
  const dispatch = useDispatch()
  const setZone = value => dispatch(setURLParam('tz', value, 'local'))
  const { data, loading, error } = useQuery(tzQuery, {
    pollInterval: 0,
    variables: { id: props.scheduleID },
  })

  let label, tz
  if (error) {
    label = 'Error: ' + (error.message || error)
  } else if (loading) {
    label = 'Fetching timezone information...'
  } else {
    tz = data.schedule.timeZone
    label = props.label ? props.label(tz) : `Show times in ${tz}`
  }

  return (
    <FormControlLabel
      control={
        <Switch
          checked={zone !== 'local'}
          onChange={e => {
            const newZone = e.target.checked ? tz : 'local'
            setZone(newZone)
            if (typeof props.onChange === 'function') {
              return props.onChange(newZone)
            }
          }}
          value={tz}
          disabled={Boolean(loading || error)}
        />
      }
      label={label}
    />
  )
}
ScheduleTZFilter.propTypes = {
  label: p.func,
  scheduleID: p.string.isRequired,

  // to be executed and returned after the zone is set in the URL
  onChange: p.func,

  // provided by connect
  zone: p.string,
  setZone: p.func,
}
